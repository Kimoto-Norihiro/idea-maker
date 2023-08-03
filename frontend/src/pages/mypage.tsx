import { NextPage } from 'next'
import React, { useEffect, useState } from 'react'
import { HeaderWithBody } from '@/components/layouts/HeaderWithBody'
import { BsPerson } from 'react-icons/bs'
import { GrFormAdd } from 'react-icons/gr'
import { TfiMore } from 'react-icons/tfi'
import axios from 'axios'
import { Theme, User } from '@/types/types'
import WorkSpace from '../components/mypage/WorkSpace';
import { useRouter } from 'next/router'

const MyPage: NextPage = () => {
	const router = useRouter()
	const [presentTheme, setPresentTheme] = useState<Theme | null>(null)
	const [currentUser, setCurrentUser] = useState<User | null>(null)

	const getUser = async () => {
		try {
			const token = localStorage.getItem('token')
			const { data } = await axios.get('http://localhost:8080/user', {
				headers: {'Authorization': `Bearer ${token}`},
				withCredentials: true,
			})
			setCurrentUser(data.data)
		} catch (err: any) {
			console.log(err)
		}
	}

	useEffect(() => {
		const token = localStorage.getItem('token')
		if (!token) router.push('/signin')
		getUser()
	}, [])

	return (
		<HeaderWithBody>
			<div className='flex h-full'>
				<div className='w-[20vw] h-full border-gray-100 border-2 border-r-0 bg-gray-50'>
					<div className='h-[4vh] flex items-center pl-4 border-gray-100 border-b-2'>
						<BsPerson/>
						<p className='ml-4'>My Themes</p>
					</div>
					<div className=''>
						<div className="flex h-[4vh] pl-3 items-center">
							<div className='hover:bg-gray-200 flex items-center p-1 rounded-md'>
								<GrFormAdd/>
							</div>
							<p className='ml-3'>add theme</p>
						</div>
						{
							currentUser?.themes?.map((theme, index) => (
								<div key={index} className="h-[4vh] hover:bg-gray-200 flex items-center justify-between px-4 group" 
									onClick={() => setPresentTheme(theme)}
								>
									<p className="">{theme.name}</p>
									<div className='invisible group-hover:visible p-1 hover:bg-gray-300 rounded-md'>
										<TfiMore/>
									</div>
								</div>
							))
						}
					</div>
				</div>
				<WorkSpace theme={presentTheme}/>
			</div>
		</HeaderWithBody>
	)
}

export default MyPage