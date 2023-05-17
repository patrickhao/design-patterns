package main

import "fmt"

// 类中有一系列相同的步骤，可以封装成模版，并实现这些步骤的统一调用方法
// 由于go接口中不允许有实现，因此这里在struct Otp中实现
type IOtp interface {
	genRandomOTP(int) string
	saveOTPCache(string)
	getMessage(string) string
	sendNotification(string) error
}

type Otp struct {
	iOtp IOtp
}

func (o *Otp) getAndSendOTP(optLength int) error {
	otp := o.iOtp.genRandomOTP(optLength)
	o.iOtp.saveOTPCache(otp)
	message := o.iOtp.getMessage(otp)
	err := o.iOtp.sendNotification(message)
	if err != nil {
		return err
	}

	return nil
}

// 各种类实现模版中的方法，并且可以统一调用
type Sms struct {
	Otp
}

func (s *Sms) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("SMS: generating ramdom otp %s\n", randomOTP)
	return randomOTP
}

func (s *Sms) saveOTPCache(otp string) {
	fmt.Printf("SMS: saving otp: %s to cache\n", otp)
}

func (s *Sms) getMessage(otp string) string {
	return "STM OTP for login is " + otp
}

func (S *Sms) sendNotification(message string) error {
	fmt.Printf("SMS: sending sms: %s\n", message)
	return nil
}

type Email struct {
	Otp
}

func (s *Email) genRandomOTP(len int) string {
	randomOTP := "1234"
	fmt.Printf("Email: generating ramdom otp %s\n", randomOTP)
	return randomOTP
}

func (s *Email) saveOTPCache(otp string) {
	fmt.Printf("Email: saving otp: %s to cache\n", otp)
}

func (s *Email) getMessage(otp string) string {
	return "Email OTP for login is " + otp
}

func (S *Email) sendNotification(message string) error {
	fmt.Printf("Email: sending sms: %s\n", message)
	return nil
}

func main() {
	smsOTP := &Sms{}
	o := Otp{
		iOtp: smsOTP,
	}
	o.getAndSendOTP(4)

	emailOTP := &Email{}
	o = Otp{
		iOtp: emailOTP,
	}
	o.getAndSendOTP(4)
}
