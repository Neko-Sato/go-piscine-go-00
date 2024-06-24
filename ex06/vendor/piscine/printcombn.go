/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcombn.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hshimizu <hshimizu@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 21:18:01 by hshimizu          #+#    #+#             */
/*   Updated: 2024/06/24 23:01:04 by hshimizu         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func printNDigitNumber(nb int, n int) {
	var buff [10]rune
	i := 0
	for ; i < n; i++ {
		buff[i] = rune('0' + nb%10)
		nb /= 10
	}
	i--
	for ; 0 <= i; i-- {
		ft.PrintRune(buff[i])
	}
}

func init_comb(s int, n int) int {
	if n <= 1 {
		return s
	}
	return init_comb(s*10+s%10+1, n-1)
}

func next_comb(n int) int {
	if n%10 == 9 {
		n /= 10
		n = next_comb(n)
		n = n*10 + n%10
		n = next_comb(n)
	} else {
		n++
	}
	return n
}

func PrintCombN(n int) {
	if n < 1 || 9 < n {
		return
	}
	isfirst := true
	i := init_comb(0, n)
	end := init_comb(10-n, n)
	for ; i < end; i = next_comb(i) {
		if isfirst {
			isfirst = false
		} else {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
		printNDigitNumber(i, n)
	}
	ft.PrintRune('\n')
}
