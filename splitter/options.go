package splitter

type splitterOpt func(*splitter)

func WithLeadingSize(size int) splitterOpt {
	return func(s *splitter) {
		s.leadingSize = size
	}
}

func WithRegularSize(size int) splitterOpt {
	return func(s *splitter) {
		s.regularSize = size
	}
}

func WithPageTmpl(tmpl string) splitterOpt {
	return func(s *splitter) {
		s.pageTmpl = tmpl
	}
}

func WithNextTmpl(tmpl string) splitterOpt {
	return func(s *splitter) {
		s.nextTmpl = tmpl
	}
}

func WithLeadingPageAddition(a string) splitterOpt {
	return func(s *splitter) {
		s.leadingPageAddition = a
	}
}
