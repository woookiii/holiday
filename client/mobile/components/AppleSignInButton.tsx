import { StyleSheet, View } from "react-native";
import {
  AppleAuthenticationButton,
  AppleAuthenticationButtonStyle,
  AppleAuthenticationButtonType,
  AppleAuthenticationScope,
  signInAsync,
} from "expo-apple-authentication";
import { useAuth } from "@/hooks/useAuth";

export default function AppleSignInButton() {
  const { signInWithAppleMutation } = useAuth();

  return (
    <View>
      <AppleAuthenticationButton
        buttonType={AppleAuthenticationButtonType.SIGN_IN}
        buttonStyle={AppleAuthenticationButtonStyle.BLACK}
        cornerRadius={5}
        style={styles.appleButton}
        onPress={async () => {
          const credential = await signInAsync({
            requestedScopes: [AppleAuthenticationScope.EMAIL],
          });
          if (credential.email) {
            signInWithAppleMutation.mutate({
              user: credential.user,
              identityToken: credential.identityToken,
              email: credential.email
            });
          }
        }}
      />
    </View>
  );
}

const styles = StyleSheet.create({
  appleButton: {
    width: 300,
    height: 44,
  },
});
