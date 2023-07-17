// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAuth } from "firebase/auth";
import firebase from 'firebase/app';

// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyAMiuKN9fFJmnQ5PETDn0RTjnhPPAaVYQU",
  authDomain: "idea-ma.firebaseapp.com",
  projectId: "idea-ma",
  storageBucket: "idea-ma.appspot.com",
  messagingSenderId: "122042054240",
  appId: "1:122042054240:web:df9f9b8cc6eaa51d3cf505",
  measurementId: "G-WC94BDXGK3"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
export const auth = getAuth(app);

