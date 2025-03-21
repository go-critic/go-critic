package checker_test

import "context"

type panel struct {
	name   string
	width  int
	height int
}

func newPanel(name string, opts ...func(c *panel)) *panel {
	cfg := &panel{
		name:   name,
		width:  0,
		height: 0,
	}
	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}

func withWidth(w int) func(c *panel) {
	return func(c *panel) {
		c.width = w
	}
}

func withHeight(h int) func(c *panel) {
	return func(c *panel) {
		c.height = h
	}
}

type exportOptions struct{}

type exportOpt func(context.Context, *exportOptions) error

func withImage(_ any, _ string) exportOpt {
	return func(_ context.Context, _ *exportOptions) error {
		return nil
	}
}

func export(_ any, _ ...exportOpt) {
}

func doSome1(w, h int) {
	_ = newPanel("hello",
		withWidth(w),
		withHeight(h),
		/*! func arg `withWidth(w)` is duplicated */
		withWidth(w),
	)

	export(nil,
		withImage(nil, ""),
		/*! func arg `withImage(nil, "")` is duplicated */
		withImage(nil, ""),
	)
}
