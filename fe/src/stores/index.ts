import { configureStore } from "@reduxjs/toolkit";
import itemSlice from "./slices/itemSlices";

const store = configureStore({
  reducer: {
    itemSlice
  },
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;

export default store;
