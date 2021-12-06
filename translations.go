// Package util
// @Author:        asus
// @Description:   $
// @File:          translations
// @Data:          2021/12/39:36
//
package util

import (
	"errors"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	"reflect"
)

type Translations struct {
	label    string
	v        *validator.Validate
	rules    *map[string]validator.Func
	rulesMsg *map[string]RegisterTrans
	trans    ut.Translator
}

type RegisterTrans struct {
	RegisterTranslationsFunc validator.RegisterTranslationsFunc
	TranslationFunc          validator.TranslationFunc
}

type Option interface {
	apply(v *Translations)
}

type optionFunc func(*Translations)

func (f optionFunc) apply(o *Translations) {
	f(o)
}

// WithLabelOption 字段转换成中文标签
func WithLabelOption(label string) Option {
	return optionFunc(func(v *Translations) {
		v.label = label
	})
}

// WithRulesOption 注册自定义验证方法
func WithRulesOption(rules *map[string]validator.Func) Option {
	return optionFunc(func(v *Translations) {
		v.rules = rules
	})
}

// WithRulesMsgOption 注册自定义错误信息
func WithRulesMsgOption(rulesMsg *map[string]RegisterTrans) Option {
	return optionFunc(func(v *Translations) {
		v.rulesMsg = rulesMsg
	})
}

// GetTranslationIns 返回参数验证的实列
func NewTranslationIns(options ...Option) *Translations {
	transIns := &Translations{
		v: validator.New(),
	}

	// 转换成中文
	transIns.SetTrans()

	// 设置内容
	for _, option := range options {
		option.apply(transIns)
	}

	// 字段转换为标签定义字段
	if transIns.label != "" {
		transIns.v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			return string(fld.Tag.Get(transIns.label))
		})
	}

	// 注册自定义验证规则
	if transIns.rules != nil {
		for key, f := range *transIns.rules {
			transIns.v.RegisterValidation(key, f)
		}
	}

	// 注册自定义错误内容
	if transIns.rulesMsg != nil {
		for key, val := range *transIns.rulesMsg {
			transIns.v.RegisterTranslation(key, transIns.trans, val.RegisterTranslationsFunc, val.TranslationFunc)
		}
	}

	return transIns
}

// SetTrans 中文翻译
func (translations *Translations) SetTrans() {
	uni := ut.New(zh.New())
	translations.trans, _ = uni.GetTranslator("zh")
	zh_translations.RegisterDefaultTranslations(translations.v, translations.trans)
}

// ValidateParam 校验参数
func (translations *Translations) ValidateParam(data interface{}) (errRes error) {
	err := translations.v.Struct(data)
	if err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			return err
		}
		str := ""
		for _, val := range errs.Translate(translations.trans) {
			str += val + ","
		}

		return errors.New(str[:len(str)-1])
	}

	return nil
}
