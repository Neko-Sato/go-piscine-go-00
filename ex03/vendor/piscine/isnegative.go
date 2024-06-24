/* ************************************************************************** */
/*                                                                            */
/*                                                        :::      ::::::::   */
/*   isnegative.go                                      :+:      :+:    :+:   */
/*                                                    +:+ +:+         +:+     */
/*   By: hshimizu <hshimizu@student.42.fr>          +#+  +:+       +#+        */
/*                                                +#+#+#+#+#+   +#+           */
/*   Created: 2024/06/24 21:18:01 by hshimizu          #+#    #+#             */
/*   Updated: 2024/06/24 21:31:23 by hshimizu         ###   ########.fr       */
/*                                                                            */
/* ************************************************************************** */

package piscine

import "ft"

func IsNegative(nb int) {
	if (nb < 0) {
		ft.PrintRune('T')
	} else {
		ft.PrintRune('F')
	}
	ft.PrintRune('\n')
}
