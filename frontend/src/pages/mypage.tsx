import { NextPage } from 'next'
import React from 'react'
import { HeaderWithBody } from '@/components/layouts/HeaderWithBody'
import { BsPerson, BsLightbulb } from 'react-icons/bs'
import { GrFormAdd } from 'react-icons/gr'
import { MdWorkspacesOutline } from 'react-icons/md'
import { PiCards } from 'react-icons/pi'
import { BiBrain } from 'react-icons/bi'
import { TfiMore } from 'react-icons/tfi'

const MyPage: NextPage = () => {
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
						<div className="h-[4vh] hover:bg-gray-200 flex items-center justify-between px-4 group">
							<p className="">Theme1</p>
							<div className='invisible group-hover:visible p-1 hover:bg-gray-300 rounded-md'>
								<TfiMore/>
							</div>
						</div>
						<div className="h-[4vh] hover:bg-gray-200 flex items-center justify-between px-4 group">
							<p className="">Theme2</p>
							<div className='invisible group-hover:visible p-1 hover:bg-gray-300 rounded-md'>
								<TfiMore/>
							</div>
						</div>
						<div className="h-[4vh] hover:bg-gray-200 flex items-center justify-between px-4 group">
							<p className="">Theme3</p>
							<div className='invisible group-hover:visible p-1 hover:bg-gray-300 rounded-md'>
								<TfiMore/>
							</div>
						</div>
					</div>
				</div>
				<div className='w-[80vw] h-full border-gray-100 border-2'>
					<div className='h-[4vh] w-[80vw] flex items-center pl-4 border-gray-100 border-b-2'>
						<MdWorkspacesOutline/>
						<p className='ml-4'>Theme1 workspace</p>
					</div>
					<div className='flex h-[88vh] w-full'>
						<div className='w-[40vw] h-[88vh] border-r-2 border-b-2 border-gray-100'>
							<div className='h-[4vh] flex items-center pl-4 border-gray-100 border-b-2'>
								<PiCards/>
								<p className='ml-4'>elements</p>
							</div>
							<div className="flex h-[4vh] pl-3 items-center">
								<div className='hover:bg-gray-200 flex items-center p-1 rounded-md'>
									<GrFormAdd/>
								</div>
								<p className='ml-3'>add element</p>
							</div>
							<div className="h-[4vh] hover:bg-gray-200 flex items-center pl-4">
								<p className="">element1</p>
							</div>
							<div className="h-[4vh] hover:bg-gray-200 flex items-center pl-4">
								<p>element2</p>
							</div>
							<div className="h-[4vh] hover:bg-gray-200 flex items-center pl-4">
								<p>element3</p>
							</div>
						</div>
						<div className='w-[40vw] h-[88vh] border-r-2 border-b-2 border-gray-100'>
							<div className='h-[4vh] flex items-center pl-4 border-gray-100 border-b-2'>
								<BsLightbulb/>
								<p className='ml-4'>Ideas</p>
							</div>
							<div className='flex'>
								<div className="flex h-[4vh] w-[20vw] pl-3 items-center">
									<div className='hover:bg-gray-200 flex items-center p-1 rounded-md'>
										<GrFormAdd/>
									</div>
									<p className='ml-3'>add idea</p>
								</div>
								<div className="flex h-[4vh] w-[20vw] pl-3 items-center">
									<div className='hover:bg-gray-200 flex items-center p-1 rounded-md'>
										<BiBrain/>
									</div>
									<p className='ml-3'>generate idea with AI</p>
								</div>
							</div>
							<div className="h-[4vh] hover:bg-gray-200 flex items-center pl-4">
								<p className="">idea1</p>
							</div>
							<div className="h-[4vh] hover:bg-gray-200 flex items-center pl-4">
								<p>idea2</p>
							</div>
							<div className="h-[4vh] hover:bg-gray-200 flex items-center pl-4">
								<p>idea3</p>
							</div>
						</div>
					</div>	
				</div>
			</div>
		</HeaderWithBody>
	)
}

export default MyPage