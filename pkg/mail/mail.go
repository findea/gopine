package mail

import (
	"fmt"
	"goweb/pkg/log"
	"goweb/pkg/util/path"
	"goweb/pkg/util/tmpl"
	"io/ioutil"
	"strings"
)

type Attach struct {
	Name, Path string
}

type Mail interface {
	SendMail(to, subject, body string, attachs ...*Attach) error
}

func SendMail(to, subject, body string, attachs ...*Attach) (err error) {
	if len(configs) < 1 {
		return loggerMail.SendMail(to, subject, body, attachs...)
	}

	var best *Config
	for _, conf := range configs {
		if strings.HasSuffix(to, conf.Key) {
			best = conf
			break
		}
	}

	var mails = make([]Mail, 0, len(configs))
	if best != nil {
		mails = append(mails, best)
	}

	for _, conf := range configs {
		if conf != best {
			mails = append(mails, conf)
		}
	}

	for _, mail := range mails {
		err = mail.SendMail(to, subject, body, attachs...)
		if err != nil {
			log.Error(err)
			continue
		}
		break
	}

	return
}

func SendMailWithTmpl(to, tmplName string, params interface{}, attachs ...*Attach) error {
	// find template path
	tmplFile := fmt.Sprintf("template/mail_%s.gohtml", tmplName)
	tmplPath, err := path.FindPath(tmplFile, 5)
	if err != nil {
		return err
	}

	// file content
	content, err := ioutil.ReadFile(tmplPath)
	if err != nil {
		return err
	}
	slices := strings.SplitN(string(content), "\n", 2)

	// subject
	subject, err := tmpl.Text(fmt.Sprintf("%s:subject", tmplPath), slices[0], params)
	if err != nil {
		return err
	}
	subject = strings.TrimSpace(subject)
	subject = strings.TrimPrefix(subject, "<!--")
	subject = strings.TrimSuffix(subject, "-->")
	subject = strings.TrimSpace(subject)

	// body
	body, err := tmpl.HTML(fmt.Sprintf("%s:body", tmplPath), slices[1], params)
	if err != nil {
		return err
	}

	return SendMail(to, subject, body, attachs...)
}
