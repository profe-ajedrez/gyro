# gyro
Decimal over a int128 implementation

```
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣀⣤⣶⣎⠭⠭⠦⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣠⠴⠚⣩⠔⣋⣥⠶⠀⠀⠀⣄⠉⠓⡄⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠁⠀⠀⠀⠀⠀⠀⠀
⠀⠀⠀⠀⠀⠀⠀⠀⢀⡠⠴⠊⠁⠀⢸⠟⣡⠾⠋⠀⣠⡶⠀⠀⢻⣇⣀⣿⡤⠤⠤⣤⣄⣤⣤⠔⠂⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⣀⣀⣤
⠀⠀⠀⠀⠀⠀⣠⢞⡭⠀⢠⣶⡀⢠⠏⡼⠋⢀⣤⡿⠋⢀⣾⢀⣾⣋⣿⠿⠒⠚⠛⠋⠉⠉⠁⠀⣀⣀⣠⣤⣶⣶⣶⣾⣿⣿⣿⣿⣿⣿
⠀⠀⠀⠀⠀⣾⣴⣏⠀⠀⠈⠻⣷⡼⠀⠀⠰⠟⠁⠀⣰⠿⣣⣾⣧⣴⡿⠶⠆⢀⣠⣤⢶⣶⣿⣿⣿⣿⣿⣿⣿⣿⣿⣿⣷⣶⣿⣿⣿⣿
⠀⠀⠀⠀⡴⠛⠉⢹⣧⠀⠀⠀⢈⣿⣄⠀⠀⢠⠴⢛⡡⠞⢉⡩⠟⠀⣠⣴⡾⠋⠉⠀⠈⣿⣿⣽⣿⣿⣿⣿⣛⣛⠿⢿⣿⣿⣿⣿⡿⠁
⠀⠀⠀⠀⡷⠒⠉⠉⠉⠁⠲⡖⠉⣿⣿⣷⣤⣴⣮⣭⠴⠒⢉⣤⡶⠛⠉⠀⣿⣴⣦⣄⠀⠻⣿⣿⣿⡿⣿⣿⣿⣿⣿⣿⣿⣿⣿⠟⠀⠀
⠀⠀⢀⡞⣡⡶⠀⠀⠒⢦⡀⠱⣤⣿⣿⠟⠿⠿⠏⠀⣠⣴⣿⣥⣤⣤⣤⣴⠟⣷⠈⠙⠳⣄⠀⠐⢽⡻⢷⣮⣽⣿⣿⣿⣿⣟⠁⣀⣀⠀
⠀⠀⣾⣾⢋⣴⡶⠂⠀⢸⡗⢳⡿⣯⠀⠀⠀⢀⣴⣞⣻⣿⡿⢚⡿⠿⣿⣯⣀⠙⢶⣖⣒⠚⠀⢠⣄⠈⢦⠈⠙⠻⣿⡟⠁⠈⠉⢀⣼⠤
⠀⠀⢿⠃⣼⠟⢀⣼⡇⣾⢃⣾⠁⠿⢃⠖⣠⡿⢿⣾⠟⢹⣧⣬⣽⡿⣫⠿⠯⢷⡈⠛⠛⡃⠀⣷⡈⣦⡈⡇⠀⠀⠈⢳⣄⡀⠀⠀⠀⠀
⠀⠀⠈⢾⣏⢀⣼⣋⡴⢛⡿⠁⢀⠞⣡⣾⣿⡀⠀⠑⠒⠀⠀⠀⠉⠋⠀⠀⠀⠀⠹⡀⠀⠹⡔⠻⣿⡏⠁⣹⣄⠀⠀⠀⠙⢿⠒⠤⣀⠀
⠀⠀⠀⠀⠈⠙⠛⠣⣶⡟⠁⡰⢃⡴⠿⣛⣿⣷⣤⠶⠦⠤⡀⠀⠀⠀⠀⠀⠀⠀⠀⣇⢠⡤⢧⢠⠟⢀⡰⠃⠘⠢⡀⠀⢱⡈⠆⢢⠀⠙
⠀⠀⠀⠀⠀⠀⠀⣰⣿⢀⠞⣡⣿⣀⣴⢿⣽⣿⡇⠀⠀⢀⣸⡀⠀⠀⠀⠀⠀⠀⠀⠘⠀⢇⠘⣿⣶⡏⠀⠀⢤⡀⠈⢢⠀⢷⡘⠈⡄⠀
⠀⠀⠀⠀⠀⠀⣴⠉⣩⢃⣼⡇⠈⢩⠁⠈⠻⡼⣧⣄⠀⢀⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⠉⢻⠉⡇⠀⠀⢠⢹⡦⡀⢣⠘⣇⠸⠇⠀
⠀⠀⠀⠀⠀⣰⢃⡜⢡⡞⠋⠻⣾⣏⠀⠀⠀⢿⣄⣨⡽⠿⢻⠿⠿⢷⣦⠀⠀⠀⠀⠀⠀⣠⣄⡸⠀⣿⣦⣄⠀⠳⣵⡌⠢⡇⢸⠁⠀⠀
⠀⠀⠀⠀⣰⣣⠋⣰⣿⡇⠀⢀⠨⢿⣄⠀⠀⠀⠉⢟⢶⣖⠋⣀⡤⠖⠋⠀⡀⠀⠀⠀⠰⡇⢸⠇⠀⣿⣿⣿⣦⡀⠙⣿⡄⢹⠸⡇⠀⠀
⠀⠀⠀⣼⡿⢁⣾⠯⣿⣧⣠⡸⣄⠀⠻⡆⠀⠀⠀⠈⠃⠛⠛⠉⠀⠀⣀⣼⠇⠀⠀⠀⠀⢹⠏⠀⢠⣿⣿⣿⣿⣿⣄⠈⢧⢸⣾⡇⠀⠀
⠀⢀⡼⠋⣠⣿⣿⡄⣿⣿⣿⡄⠈⢻⡗⣻⣶⣄⠀⠀⠀⢰⣤⣴⠾⠛⠉⠁⠀⢠⡒⠉⢷⡿⠀⠀⣼⣯⣟⢿⣿⣿⣿⣷⣼⣿⣿⠀⠀⢀
⢀⠞⢁⣼⣿⣿⣿⣣⣿⣿⣿⣇⠀⠀⡟⢻⣿⣿⣷⣀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢣⣴⢻⠃⠀⢀⣿⣿⣿⣧⠉⢻⣿⣿⣿⠃⠁⠀⢀⣌
⠏⣠⣿⣯⣿⣿⣿⣿⣿⣿⣿⣿⡀⠀⣇⢸⢿⣿⣿⣿⣷⣤⡒⢆⠀⠀⢀⣠⣴⣾⡿⢁⡎⠀⠀⣼⠿⣿⣿⣿⣇⢸⣿⣿⡇⠀⠀⢠⡿⠃
⣾⣿⠿⠟⣛⣿⣿⣿⣿⣿⣿⣿⣇⠀⢻⣾⢸⡘⣿⣿⣿⣿⣿⣬⣷⣾⣿⣿⣿⠟⠀⠈⠀⠀⢠⡿⠀⢻⣿⣟⣿⢸⣿⣿⠃⠀⣰⡿⠁⠀
⠀⣠⠴⠿⠿⠿⠿⠿⣿⣿⣿⣿⣿⡄⠈⣿⡀⢷⣿⣿⡟⣿⣿⣿⣧⣿⣿⡿⠋⠀⠀⠀⠀⢠⣿⡇⢠⣿⢟⣾⠟⣸⣿⡿⢀⣼⡿⠃⠀⠀
⠉⠀⠀⠉⣉⠽⠋⠉⢀⡤⢈⠏⢻⣧⠀⢹⣇⠘⣿⣿⣷⣏⡿⣿⣸⣿⡿⠁⠀⠀⠀⠀⢠⡿⢻⣥⣿⣿⣿⡟⢰⣿⡟⢡⡾⡽⠁⠀⠀⠐
⠙⢿⣿⣞⠙⢷⣶⣿⡿⠖⠉⢀⡄⠹⡄⠈⢟⡆⢸⣟⢿⣿⣿⣿⣿⣽⣇⠀⠀⠀⠀⠀⠘⠃⣾⣷⣿⣿⡟⢠⡿⠋⡰⢋⡞⢁⡠⠂⢈⡆
```


## Benchmarks

```
go test -benchmem -benchtime=4s -memprofile=mem.out -run=^$ -bench ./...     

goos: linux
goarch: amd64
pkg: github.com/profe-ajedrez/gyro
cpu: Intel(R) Core(TM) i7-10700 CPU @ 2.90GHz
BenchmarkNew/Zero_coeff_and_exp-16         	1000000000	         1.584 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew/Positive_coeff_and_exp-16     	1000000000	         4.080 ns/op	       0 B/op	       0 allocs/op
BenchmarkNew/Negative_coeff_and_exp-16     	1000000000	         1.959 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroString/Zero_coeff_and_exp-16  	1000000000	         2.397 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroString/Positive_coeff_and_zero_exp-16         	222617127	        21.57 ns/op	       5 B/op	       1 allocs/op
BenchmarkGyroString/Negative_coeff_and_zero_exp-16         	89079247	        52.99 ns/op	      16 B/op	       2 allocs/op
BenchmarkGyroString/Positive_coeff_and_positive_exp-16     	189682452	        25.33 ns/op	       8 B/op	       1 allocs/op
BenchmarkGyroString/Negative_coeff_and_positive_exp-16     	72444318	        65.74 ns/op	      32 B/op	       2 allocs/op
BenchmarkGyroString/Positive_coeff_and_negative_exp-16     	71666110	        65.80 ns/op	      10 B/op	       2 allocs/op
BenchmarkGyroString/Negative_coeff_and_negative_exp-16     	53724480	        88.77 ns/op	      16 B/op	       3 allocs/op
BenchmarkNewFromString/Valid_integer-16                    	333265204	        14.44 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewFromString/Valid_negative_integer-16           	316131117	        15.10 ns/op	       0 B/op	       0 allocs/op
BenchmarkNewFromString/Valid_decimal_with_positive_exponent-16         	30444026	       156.9 ns/op	      48 B/op	       3 allocs/op
BenchmarkNewFromString/Valid_decimal_with_negative_exponent-16         	27509752	       174.0 ns/op	      48 B/op	       3 allocs/op
BenchmarkGyroAdd/Zero_and_zero-16                                      	1000000000	         2.404 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroAdd/Positive_and_zero-16                                  	1000000000	         2.410 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroAdd/Negative_and_zero-16                                  	1000000000	         2.396 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroAdd/Positive_and_positive-16                              	1000000000	         2.398 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroAdd/Negative_and_negative-16                              	1000000000	         2.405 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroAdd/Positive_and_negative-16                              	1000000000	         2.410 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroAdd/Negative_and_positive-16                              	1000000000	         2.402 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroAdd/Different_exponents-16                                	1000000000	         2.403 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroAdd/Large_values-16                                       	1000000000	         2.405 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Zero_minus_zero-16                                    	1000000000	         2.399 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Positive_minus_zero-16                                	1000000000	         2.401 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Negative_minus_zero-16                                	1000000000	         2.403 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Positive_minus_positive_(larger_-_smaller)-16         	1000000000	         2.401 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Positive_minus_positive_(smaller_-_larger)-16         	1000000000	         2.394 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Negative_minus_negative_(larger_-_smaller)-16         	1000000000	         2.397 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Negative_minus_negative_(smaller_-_larger)-16         	1000000000	         2.396 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Positive_minus_negative-16                            	1000000000	         2.408 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Negative_minus_positive-16                            	1000000000	         2.397 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Different_exponents-16                                	1000000000	         2.405 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyroSub/Large_values-16                                       	1000000000	         2.396 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[_0]_-10.10112212_/_2.304_w_16_exp-16                     	122027338	        39.27 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[_1]_0.10_/_0.3_w_3_exp-16                                	195313158	        24.54 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[_2]_10_/_3_w_1_exp-16                                    	222249886	        21.61 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[_3]_10_/_2_w_0_exp-16                                    	222131586	        21.62 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[_4]_10_/_-3_w_0_exp-16                                   	208164154	        22.99 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[_5]_10_/_3_w_0_exp-16                                    	221608048	        21.60 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[_6]_-10_/_-2_w_0_exp-16                                  	206334907	        23.24 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[_7]_10_/_-2_w_0_exp-16                                   	207407497	        23.20 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[_8]_-10_/_2_w_0_exp-16                                   	207956461	        23.08 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[_9]_1_/_10_w_0_exp-16                                    	207995438	        23.07 ns/op	       0 B/op	       0 allocs/op
BenchmarkDiv/[10]_0_/_10_w_0_exp-16                                    	221380321	        21.72 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyro_Round/<2_decimals_up>-16                                 	1000000000	         0.0000011 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyro_Round/<4_decimals_up>-16                                 	1000000000	         0.0000019 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyro_Round/<3_decimals_up>-16                                 	1000000000	         0.0000022 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyro_Round/<2_decimals_up>#01-16                              	1000000000	         0.0000017 ns/op	       0 B/op	       0 allocs/op
BenchmarkGyro_Round/<2_decimals_down>-16                               	1000000000	         0.0000013 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	github.com/profe-ajedrez/gyro	199.705s
```

