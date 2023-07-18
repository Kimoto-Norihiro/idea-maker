import React, { useState } from 'react'
import Link from 'next/link'
import type { NextPage } from 'next'
import { set, useForm } from 'react-hook-form'
import * as yup from 'yup'
import { yupResolver } from '@hookform/resolvers/yup'
import { User } from '@/types/types'
import { useRouter } from 'next/router'
import axios from 'axios'
import { InputWithError } from '../components/InputWithError';
import { HeaderWithBody } from '../components/layouts/HeaderWithBody';
import { signInWithEmailAndPassword } from 'firebase/auth'
import { auth } from '@/firebase/firebase'

const signInSchema = yup.object().shape({
  email: yup.string().email('invalid email').required('required input'),
  password: yup.string().required('required input').min(8, 'at least 8 characters'),
})

type signInFormValues = Omit<User, 'name'>

const SignIn: NextPage = () => {
  const router = useRouter()
  const [err, setErr] = useState('')

  const { register, handleSubmit, formState: { errors } } = useForm<signInFormValues>({
    resolver: yupResolver(signInSchema)
  });

  const firebaseSignIn = async (data: signInFormValues) => {
    try {
      const response = await signInWithEmailAndPassword(auth, data.email, data.password)
      const token = await response.user.getIdToken()
      localStorage.setItem('token', token)
      console.log("setToken", token)
      return response.user
    } catch(err: any) {
      setErr('error')
      console.log("firebase SignIn", err.message)
      return null
    }
  }

  const submit = () => {
    handleSubmit(async (data) => {
      const user = await firebaseSignIn(data)
      if (user) {
        router.push('/')
      }
    }, () => {
      console.log('error')
      setErr('error')
    })()
  }

  return (
    <HeaderWithBody>
      <div className='flex flex-col bg-white justify-center p-10 align-center'>
        <div className='flex justify-center text-3xl'>
          <h1>Sign In</h1>
        </div>
        <div className='flex p-5 justify-center bg-white'>
          <form
            className='w-96 p-3' 
            onSubmit={(e) => {
            e.preventDefault()
            submit()
          }}>
            <InputWithError 
              name='email'
              register={register}
              errors={errors}
            />
            <InputWithError 
              name='password'
              register={register}
              errors={errors}
            />
            {
              err ? <div className='text-red-800 text-sm'>{err}</div> : <div className='h-5'></div>
            }
            <div>
              if you don`t have account ... <Link href='signup' className='text-blue-700 underline'>signup</Link>
            </div>
            <div className='flex justify-center mt-5'>
              <div className='bg-gray-200 py-1 px-3 rounded-lg'>
                <button type='submit'>sign in</button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </HeaderWithBody>
  )
}

export default SignIn