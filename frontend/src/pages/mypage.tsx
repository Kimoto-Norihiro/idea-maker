import { NextPage } from 'next'
import React, { useEffect, useState } from 'react'
import { HeaderWithBody } from '@/components/layouts/HeaderWithBody'
import { BsPerson } from 'react-icons/bs'
import { GrFormAdd } from 'react-icons/gr'
import { TfiMore } from 'react-icons/tfi'
import axios from 'axios'
import { Theme, User } from '@/types/types'
import WorkSpace from '../components/mypage/WorkSpace';

const MyPage: NextPage = () => {
	const [presentThemeId, setPresentThemeId] = useState<string | null>(null)
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

	const createTheme = async () => {
		try {
			const token = localStorage.getItem('token')
			console.log("create theme token", token)
			const { data } = await axios.post('http://localhost:8080/theme', undefined, {
				headers: {'Authorization': `Bearer ${token}`},
				withCredentials: true,
			})
			console.log(data)
			await getUser()
		} catch (err: any) {
			console.log("create theme:", err)
		}
	}

	useEffect(() => {
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
							<div 
								className='hover:bg-gray-200 flex items-center p-1 rounded-md'
								onClick={createTheme}	
							>
								<GrFormAdd/>
							</div>
							<p className='ml-3'>add theme</p>
						</div>
						{
							currentUser?.themes?.map((theme, index) => (
								<ThemeCard key={index} theme={theme} getUser={getUser} setPresentThemeId={setPresentThemeId} />
							))
						}
					</div>
				</div>
				<WorkSpace themeId={presentThemeId}/>
			</div>
		</HeaderWithBody>
	)
}

export default MyPage

type ThemeCardProps = {
	theme: Theme
	getUser: () => Promise<void>
	setPresentThemeId: React.Dispatch<React.SetStateAction<string | null>>
}

const ThemeCard = ({ theme, getUser, setPresentThemeId }: ThemeCardProps) => {
	const [isEdit, setIsEdit] = useState(false)
	const [name, setName] = useState(theme.name)

	const editTheme = async () => {
		const newTheme = {
			...theme,
			name,
		}
		
		try {
			const token = localStorage.getItem('token')
			const { data } = await axios.put('http://localhost:8080/theme', newTheme ,{
				headers: {'Authorization': `Bearer ${token}`},
				withCredentials: true,
			})
			await getUser()
			setIsEdit(false)
		} catch (err: any) {
			console.log(err)
		}
	}

	const onEnterDown = async (e: React.KeyboardEvent) => {
		if (e.key == "Enter") editTheme()
	}

	useEffect(() => {
		if (!isEdit) setName(theme.name)
	}, [isEdit])

	return (
		<div className="">
			{
				isEdit ? (
					<div className="h-[4vh] w-full flex items-center px-4">
						<input 
							className='rounded-sm px-1' 
							value={name} 
							onChange={(e) => setName(e.target.value)}
							onKeyDown={onEnterDown}
							onBlur={() => {
								setIsEdit(false)
							}}
						/>
					</div>
					
				) : (
					<div 
						className="h-[4vh] w-full hover:bg-gray-100 flex items-center justify-between px-4 group" 
						onClick={() => setPresentThemeId(theme.ID)}
					>
						<p 
							className='hover:bg-gray-300 rounded-sm px-1'
							onClick={() => setIsEdit(true)}
						>
							{theme.name}
						</p>
						<div 
							className='invisible group-hover:visible p-1 hover:bg-gray-300 rounded-md'
							onClick={(e) => {
								e.stopPropagation()
								console.log('more')
							}}
						>
							<TfiMore/>
						</div>
					</div>
				)
			}
		</div>
	)
}