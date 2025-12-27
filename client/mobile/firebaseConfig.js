// Import the functions you need from the SDKs you need
import { initializeApp } from "firebase/app";
import { getAnalytics } from "firebase/analytics";
import {getAuth} from "firebase/auth"
// TODO: Add SDKs for Firebase products that you want to use
// https://firebase.google.com/docs/web/setup#available-libraries

// Your web app's Firebase configuration
// For Firebase JS SDK v7.20.0 and later, measurementId is optional
const firebaseConfig = {
  apiKey: "AIzaSyDcIp79gQDT0b6TZCrt4_ZwZlxuLPmyfDk",
  authDomain: "holiday-d2992.firebaseapp.com",
  projectId: "holiday-d2992",
  storageBucket: "holiday-d2992.firebasestorage.app",
  messagingSenderId: "408693957630",
  appId: "1:408693957630:web:728cfd48104298ee3816af",
  measurementId: "G-VYXJY27KT9"
};

// Initialize Firebase
const app = initializeApp(firebaseConfig);
const analytics = getAnalytics(app);
const auth = getAuth(app);