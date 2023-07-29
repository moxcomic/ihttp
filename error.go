package ihttp

func (self *IHttp) WithError(f func(err error)) (this *IHttp) {
	this = self

	self.errorHandler = f

	return
}

func (self *IHttp) doErrorHandler() {
	if f := self.errorHandler; f != nil && self.err != nil {
		self.onceError.Do(func() {
			f(self.err)
		})
	}
}
