export interface Token {
  accessToken: string;
  refreshToken: string;
}

export interface emailSignReq {
  email: string;
  password: string;
}

export interface verifyEmailOTPReq {
  otp: string;
  verificationId: string;
}

export interface verifyEmailOTPResp {
  emailVerified: boolean;
  sessionId?: string;
}

export interface verifySMSOTPReq {
  otp: string;
  verificationId: string;
  sessionId?: string;
}

export interface verifySMSOTPResp {
  phoneNumberVerified: boolean;
  accessToken?: string;
}

export interface EmailLoginResp {
  emailVerified: boolean;
  phoneNumberVerified: boolean;
  id: string;
  sessionId?: string;
  accessToken?: string;
}

export interface signInWithAppleReq {
  identityToken: string | null;
  user: string;
  email: string | null;
}

export interface signInWithAppleResp {
  phoneNumberVerified: boolean;
  sessionId?: string;
  accessToken?: string;
}
