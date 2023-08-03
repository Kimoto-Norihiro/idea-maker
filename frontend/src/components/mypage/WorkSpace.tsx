import React, { useState, useEffect } from 'react'
import { BsLightbulb } from 'react-icons/bs'
import { GrFormAdd } from 'react-icons/gr'
import { MdWorkspacesOutline } from 'react-icons/md'
import { PiCards } from 'react-icons/pi'
import { BiBrain } from 'react-icons/bi'
import { Theme } from '@/types/types'
import axios from 'axios'

type Props = {
	theme: Theme | null
}

export const WorkSpace = ({ theme }: Props) => {
	if (!theme) return (
		<div className='w-[80vw] h-full border-gray-100 border-2 items-center flex justify-center text-gray-500'>
			Theme is not selected
		</div>
	)

	const { elements, ideas } = theme

	return (
		<div className='w-[80vw] h-full border-gray-100 border-2'>
			<div className='h-[4vh] w-[80vw] flex items-center pl-4 border-gray-100 border-b-2'>
				<MdWorkspacesOutline/>
				<p className='ml-4'>{theme.name} workspace</p>
			</div>
			<div className='flex h-[88vh] w-full'>
				<div className='w-[40vw] h-[88vh] border-r-2 border-b-2 border-gray-100'>
					<div className='h-[4vh] flex items-center pl-4 border-gray-100 border-b-2'>
						<PiCards/>
						<p className='ml-4'>elements</p>
					</div>
					<div className="flex h-[4vh] pl-3 items-center">
						<div 
							className='hover:bg-gray-200 flex items-center p-1 rounded-md'
							onClick={() => console.log('add element')}
						>
							<GrFormAdd/>
						</div>
						<p className='ml-3'>add element</p>
					</div>
					{
						!elements || elements.length === 0 ? (
							<div className="h-[4vh] hover:bg-gray-200 flex items-center pl-4">
								<p className="">element is nothing</p>
							</div>
						) : (
							elements?.map((element, index) => (
								<div key={index} className="h-[4vh] hover:bg-gray-200 flex items-center pl-4">
									<p>{element.name}</p>
								</div>
							))
						)
					}
				</div>
				<div className='w-[40vw] h-[88vh] border-r-2 border-b-2 border-gray-100'>
					<div className='h-[4vh] flex items-center pl-4 border-gray-100 border-b-2'>
						<BsLightbulb/>
						<p className='ml-4'>Ideas</p>
					</div>
					<div className='flex'>
						<div className="flex h-[4vh] w-[20vw] pl-3 items-center">
							<div 
								className='hover:bg-gray-200 flex items-center p-1 rounded-md'
								onClick={() => console.log('add idea')}
							>
								<GrFormAdd/>
							</div>
							<p className='ml-3'>add idea</p>
						</div>
						<div className="flex h-[4vh] w-[20vw] pl-3 items-center">
							<div 
								className='hover:bg-gray-200 flex items-center p-1 rounded-md'
								onClick={() => console.log('generate idea')}
							>
								<BiBrain/>
							</div>
							<p className='ml-3'>generate idea with AI</p>
						</div>
					</div>
					{
						!ideas || ideas.length === 0 ? (
							<div className="h-[4vh] hover:bg-gray-200 flex items-center pl-4">
								<p className="">idea is nothing</p>
							</div>
						) : (
							ideas?.map((idea, index) => (
								<div key={index} className="h-[4vh] hover:bg-gray-200 flex items-center pl-4">
									<p>{idea.name}</p>
								</div>
							))
						)
					}
				</div>
			</div>	
		</div>
	)
}

export default WorkSpace