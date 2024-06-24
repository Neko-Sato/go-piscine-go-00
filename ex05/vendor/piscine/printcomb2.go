/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcomb2.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hshimizu <hshimizu@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 21:18:01 by hshimizu          #+#    #+#             */
/*   Updated: 2024/06/24 22:10:34 by hshimizu         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func printTwoDigitNumber(nb int) {

	ft.PrintRune(rune('0' + (nb/10)%10))
	ft.PrintRune(rune('0' + nb%10))
}

func PrintComb2() {
	isfirst := true
	for a := 0; a < 100; a++ {
		for b := a + 1; b < 100; b++ {
			if isfirst {
				isfirst = false
			} else {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
			printTwoDigitNumber(a)
			ft.PrintRune(' ')
			printTwoDigitNumber(b)
		}
	}
	ft.PrintRune('\n')
}
