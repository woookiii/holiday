import { deleteItemAsync, getItemAsync, setItemAsync } from "expo-secure-store";

async function saveSecureStore(key:string, value: string) {
  await setItemAsync(key, value);
}

async function getSecureStore(key: string) {
  const storedDate = await getItemAsync(key);
  return storedDate;
}

async function deleteSecureStore(key: string) {
  await deleteItemAsync(key);
}


export {saveSecureStore, getSecureStore, deleteSecureStore }