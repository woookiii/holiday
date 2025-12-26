import {SafeAreaView} from "react-native-safe-area-context";
import { StyleSheet, Text, View} from "react-native";
import CustomButton from "@/components/CustomButton";
import {router} from "expo-router";

export default function AuthScreen() {

  return <SafeAreaView style={styles.container}>
    <View style={styles.buttonContainer}>
      <Text style={styles.beach}>🏖️</Text>
      <CustomButton
        label={"Start with your Phone number"}
        onPress={()=> router.push('/auth/phonenumber')}
      />
    </View>
  </SafeAreaView>;
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "white"
  },
  buttonContainer: {
    flex: 1,
    width: "100%",
    alignItems: "center",
    paddingHorizontal: 40,
    marginTop: 140,
    gap: 160,
  },
  beach: {
    fontSize: 100
  }
});
