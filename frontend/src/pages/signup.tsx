import React, { useState } from 'react'
import type { NextPage } from 'next'
import Link from 'next/link'
import { set, useForm } from 'react-hook-form'
import * as yup from 'yup'
import { yupResolver } from '@hookform/resolvers/yup'
import { User } from '@/types/types'
import { useRouter } from 'next/router'
import axios from 'axios'
import { InputWithError } from '@/components/InputWithError'
import { HeaderWithBody } from '../components/layouts/HeaderWithBody';
import { auth } from '@/firebase/firebase'
import { createUserWithEmailAndPassword } from 'firebase/auth'

const signUpSchema = yup.object().shape({
  name: yup.string().required('required input'),
  email: yup.string().email('invalid email').required('required input'),
  password: yup.string().required('required input').min(8, 'at least 8 characters'),
})

export type SignUpFormValues = Omit<User, 'uid'>

const SignUp: NextPage = () => {
  const [err, setErr] = useState('')
  const router = useRouter()
  const { register, handleSubmit, formState: { errors } } = useForm<SignUpFormValues>({
    resolver: yupResolver(signUpSchema)
  })

  const createUser = async (user: SignUpFormValues) => {
    try {
      const token = localStorage.getItem('token')
      console.log("getToken", token)
      const response = await axios.post('http://localhost:8080/v1/signup', user, {
        headers: {'Authorization': `Bearer ${token}`},
        withCredentials: true,
      })
      return response
    } catch (err: any) {
      setErr(err.message)
      console.log(err)
      return null
    }
  }

  const createFirebaseUser = async (email: string, password: string) => {
    try  {
      const { user } = await createUserWithEmailAndPassword(auth, email, password)
      const token = await user.getIdToken()
      localStorage.setItem('token', token)
      console.log("setToken", token)
      return user

    } catch (err: any) {
      setErr(err.message)
      return null
    }
  }

  const submit = () => {
    handleSubmit(async (data) => {
      const user = await createFirebaseUser(data.email, data.password)
      if (user) {
        const res = await createUser(data)
        if (res) {
          router.push('/signin')
        }
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
          <h1>Sign Up</h1>
        </div>
        <div className='flex p-5 justify-center'>
          <form 
            className='w-96 p-3' 
            onSubmit={(e) => {
              e.preventDefault()
              submit()
            }}>
            <InputWithError 
              name='name'
              register={register}
              errors={errors}
            />
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
              if you have account ... <Link href='signin' className='text-blue-700 underline'>signin</Link>
            </div>
            <div className='flex justify-center mt-5'>
              <div className='bg-gray-200 py-1 px-3 rounded-lg'>
                <button type='submit'>sign up</button>
              </div>
            </div>
          </form>
        </div>
      </div>
    </HeaderWithBody>
  )
}

export default SignUp