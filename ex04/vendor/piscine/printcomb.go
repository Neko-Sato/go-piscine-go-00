/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printcomb.go                                       :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hshimizu <hshimizu@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 21:18:01 by hshimizu          #+#    #+#             */
/*   Updated: 2024/06/24 22:08:45 by hshimizu         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintComb() {
	isfirst := true
	for a := 0; a < 10; a++ {
		for b := a + 1; b < 10; b++ {
			for c := b + 1; c < 10; c++ {
				if isfirst {
					isfirst = false
				} else {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
				ft.PrintRune(rune('0' + a))
				ft.PrintRune(rune('0' + b))
				ft.PrintRune(rune('0' + c))
			}
		}
	}
	ft.PrintRune('\n')
}
