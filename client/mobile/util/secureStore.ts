import { deleteItemAsync, getItemAsync, setItemAsync } from "expo-secure-store";

async function saveSecureStore(key: string, value: string) {
  await setItemAsync(key, value);
}

async function getSecureStore(key: string) {
  return await getItemAsync(key);
}

async function deleteSecureStore(key: string) {
  await deleteItemAsync(key);
}


export { saveSecureStore, getSecureStore, deleteSecureStore };