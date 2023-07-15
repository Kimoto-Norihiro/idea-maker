import React from 'react'
import type { NextPage } from 'next'
import Link from 'next/link'
import { HeaderWithBody } from '../components/layouts/HeaderWithBody';

export const Home: NextPage = () => {
  return (
    <HeaderWithBody>
       <div className='flex justify-center p-10'>
        <Link href='signin'>
          <p className='text-7xl'>Idea Maker</p>
        </Link>
      </div>
    </HeaderWithBody>
  )
}

export default Home
