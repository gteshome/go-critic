package checker_test

func simpleIf() {
	if true {
	}
}

func balanced() {
	if true {
		if true {
			println("1")
		}
	} else {
		if false {
			println("2")
		}
	}
}

func properElseIf() {
	var cond1, cond2 bool

	if cond1 {
	} else if cond2 {
		if cond2 {
		} else if cond1 {
		}
	}

	if cond1 {
	} else if cond2 {
	} else if cond1 {
	}
}

func complexElseBody() {
	if true {
	} else {
		// Body of more than 1 statement.
		if false {
		}
		if true {
		}
	}
}
