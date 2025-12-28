import { PhoneAuthProvider, signInWithCredential, signInWithPhoneNumber } from "firebase/auth";
import { FIREBASE_AUTH, RECAPTCHA } from "@/firebaseConfig";
import { getSecureStore } from "@/util/secureStore";
import { axiosInstance } from "@/api/axios";
import { ConfirmationResult, UserCredential } from "@firebase/auth";
import { Token } from "@/types";
import { Profiler } from "node:inspector";
import Profile = module

async function requestOtpToFirebase(phoneNumber: string): Promise<ConfirmationResult> {
  return await signInWithPhoneNumber(FIREBASE_AUTH, phoneNumber, RECAPTCHA);
}

async function postPhoneOtpToFirebase(otp: string): Promise<UserCredential> {
  const verificationId = await getSecureStore("verificationId");
  if (!verificationId) throw new Error("No verificationId found");
  const credential = PhoneAuthProvider.credential(verificationId, otp);
  return await signInWithCredential(FIREBASE_AUTH, credential);
}

async function postFirebaseTokenToServer(firebaseToken: string): Promise<Token> {
  const { data } = await axiosInstance.post("/auth/firebase-token",firebaseToken)
  return data
}

async function refreshAccessToken() {
  try {
    const res = await axiosInstance.post("/auth/refresh-token");
    return res.data;
  } catch (err: any) {
    const message = err.response?.data?.message || "Failed to refresh access token";
    throw new Error(message);
  }
}

async function getMe() {
  const { data } = await axiosInstance.get("/auth/me");

  return data;
}


export { requestOtpToFirebase, postPhoneOtpToFirebase, refreshAccessToken, postFirebaseTokenToServer, getMe };