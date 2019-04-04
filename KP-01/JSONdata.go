package main

func GetData() []string {

	JSONdata := []string{`
{ 
"Expense":[
	129,
	44,
	394,
	253,
	147
	]
}
`,
`		{
			"Expense":[-100, 40000, -6500, -230, -18, 34500, -450, 13000, -100, 5000]
		}`,
		`		{
			"Expense":[
				20,
			-20,
			3
		]
		}`,
		`		{
			"Expense":[
				90,
			-20,
			30,
-120,
-12
		]
		}`,
}

	return JSONdata
}
