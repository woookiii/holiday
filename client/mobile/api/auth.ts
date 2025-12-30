import { getSecureStore } from "@/util/secureStore";
import { axiosInstance } from "@/api/axios";
import { Token } from "@/types";
import {
  FirebaseAuthTypes,
  getAuth,
  PhoneAuthProvider, signInWithCredential,
  signInWithPhoneNumber
} from "@react-native-firebase/auth";

const authInstance = getAuth();

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


export { requestSmsOtpToFirebase, postSmsOtpToFirebase, refreshAccessToken, postFirebaseTokenToServer, getMe };