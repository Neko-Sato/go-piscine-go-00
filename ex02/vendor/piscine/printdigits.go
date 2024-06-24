/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   printdigits.go                                     :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hshimizu <hshimizu@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 21:18:01 by hshimizu          #+#    #+#             */
/*   Updated: 2024/06/24 21:27:11 by hshimizu         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func PrintDigits() {
	for c := '0'; c <= '9'; c++ {
		ft.PrintRune(c)
	}
	ft.PrintRune('\n')
}
