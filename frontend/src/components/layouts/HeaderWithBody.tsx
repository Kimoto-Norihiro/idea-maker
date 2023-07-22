import React, { ReactElement } from 'react'
import { Header } from '../parts/Header';

type Props = {
	children: ReactElement
}

export const HeaderWithBody = ({children}: Props) => {
	return (
		<div className='h-[100vh]'>
			<Header />
			<div className='h-[92vh]'>
				{ children }
			</div>
		</div>
	)
}