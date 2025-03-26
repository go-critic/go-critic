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

func withDefaultWidth() func(c *panel) {
	return withWidth(100)
}

var defaultWidth = withDefaultWidth()

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

type sourceOptions struct{}

type sourceOpt = func(context.Context, *sourceOptions) error

func withGroup(_ string) sourceOpt {
	return func(_ context.Context, _ *sourceOptions) error {
		return nil
	}
}

func withKey(_ string) sourceOpt {
	return func(_ context.Context, _ *sourceOptions) error {
		return nil
	}
}

func defaultOpt(_ context.Context, _ *sourceOptions) error {
	return nil
}

func newSource(_ any, opts ...sourceOpt) (any, error) {
	ctx := context.Background()
	o := &sourceOptions{}
	for _, opt := range opts {
		err := opt(ctx, o)
		if err != nil {
			return nil, err
		}
	}
	return nil, nil
}

func doSome1(w, h int) {
	_ = newPanel("hello",
		withWidth(w),
		withHeight(h),
		/*! function argument `withWidth(w)` is duplicated */
		withWidth(w),
	)

	export(nil,
		withImage(nil, ""),
		/*! function argument `withImage(nil, "")` is duplicated */
		withImage(nil, ""),
	)

	_ = newPanel("case2",
		withHeight(h),
		withDefaultWidth(),
		/*! function argument `withDefaultWidth()` is duplicated */
		withDefaultWidth(),
		defaultWidth,
		/*! function argument `defaultWidth` is duplicated */
		defaultWidth,
	)

	_ = newPanel("case3",
		withHeight(h),
		defaultWidth,
		/*! function argument `defaultWidth` is duplicated */
		defaultWidth,
	)

	_, _ = newSource("case4",
		withGroup(""),
		withKey("key"),
		defaultOpt,
		/*! function argument `defaultOpt` is duplicated */
		defaultOpt,
	)
}
