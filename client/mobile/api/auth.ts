import { getSecureStore } from "@/util/secureStore";
import { axiosInstance } from "@/api/axios";
import { Token } from "@/types";
import {
  FirebaseAuthTypes,
  getAuth,
  PhoneAuthProvider, signInWithCredential,
  signInWithPhoneNumber
} from "@react-native-firebase/auth";
import { getApp } from "@react-native-firebase/app";

const appInstance = getApp();
const authInstance = getAuth(appInstance);

async function requestSmsOtpToFirebase(phoneNumber: string): Promise<FirebaseAuthTypes.ConfirmationResult> {
  return await signInWithPhoneNumber(authInstance, phoneNumber);
}

async function postSmsOtpToFirebase(otp: string): Promise<FirebaseAuthTypes.UserCredential> {
  const verificationId = await getSecureStore("verificationId");
  if (!verificationId) throw new Error("No verificationId found");
  const credential = PhoneAuthProvider.credential(verificationId, otp);
  return signInWithCredential(authInstance, credential);
}

async function postFirebaseTokenToServer(firebaseToken: string): Promise<Token> {
  const { data } = await axiosInstance.post("/auth/firebase-token", firebaseToken);
  return data;
}


async function getMe() {
  const { data } = await axiosInstance.get("/auth/me");

  return data;
}


export { requestSmsOtpToFirebase, postSmsOtpToFirebase, postFirebaseTokenToServer, getMe };