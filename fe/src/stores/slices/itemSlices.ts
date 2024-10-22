import { createSlice } from '@reduxjs/toolkit';
import { Dispatch } from 'redux';
import axios from 'axios';
import { BASE_URL } from '@/constants/api.contant';
import { ResponseMessageOnly, ResponseReportResult } from '@/@types/response';
import { PenerimaanBarang, PengeluaranBarang } from '@/@types/request';
import { handleError } from '../utils/errorHandler';

interface AuthState {
    //penerimaan
    isPostPenerimaanBarangSuccess: boolean
    isPostPenerimaanBarangError: boolean
    isPostPenerimaanBarangLoading: boolean
    isPostPenerimaanBarangMsg: string | null
    //pengeluaran
    isPostPengeluaranBarangSuccess: boolean
    isPostPengeluaranBarangError: boolean
    isPostPengeluaranBarangLoading: boolean
    isPostPengeluaranBarangMsg: string | null
    //data Table
    dataBarang: ResponseReportResult[]
    dataBarangLoading: boolean
}

const initialState: AuthState = {
    //penerimaan
    isPostPenerimaanBarangSuccess: false,
    isPostPenerimaanBarangError: false,
    isPostPenerimaanBarangLoading: true,
    isPostPenerimaanBarangMsg: null,
    //pengeluaran
    isPostPengeluaranBarangSuccess: false,
    isPostPengeluaranBarangError: false,
    isPostPengeluaranBarangLoading: true,
    isPostPengeluaranBarangMsg: null,
    // data table
    dataBarang: [],
    dataBarangLoading: true,
};

export const itemSlice = createSlice({
    name: 'item/slice',
    initialState,
    reducers: {
        //penerimaan
        setIsPostPenerimaanBarangSuccess: (state, action) => { state.isPostPenerimaanBarangSuccess = action.payload; },
        setIsPostPenerimaanBarangError: (state, action) => { state.isPostPenerimaanBarangError = action.payload; },
        setIsPostPenerimaanBarangLoading: (state, action) => { state.isPostPenerimaanBarangLoading = action.payload; },
        setIsPostPenerimaanBarangMsg: (state, action) => { state.isPostPenerimaanBarangMsg = action.payload; },
        //pengeluaran
        setIsPostPengeluaranBarangSuccess: (state, action) => { state.isPostPengeluaranBarangSuccess = action.payload; },
        setIsPostPengeluaranBarangError: (state, action) => { state.isPostPengeluaranBarangError = action.payload; },
        setIsPostPengeluaranBarangLoading: (state, action) => { state.isPostPengeluaranBarangLoading = action.payload; },
        setIsPostPengeluaranBarangMsg: (state, action) => { state.isPostPengeluaranBarangMsg = action.payload; },
        // data table
        setIsdataBarang: (state, action) => { state.dataBarang = action.payload; },
        setIsdataBarangLoading: (state, action) => { state.dataBarangLoading = action.payload; },
    },
});

export const {
    //penerimaan
    setIsPostPenerimaanBarangSuccess,
    setIsPostPenerimaanBarangError,
    setIsPostPenerimaanBarangLoading,
    setIsPostPenerimaanBarangMsg,
    //pengeluaran
    setIsPostPengeluaranBarangSuccess,
    setIsPostPengeluaranBarangError,
    setIsPostPengeluaranBarangLoading,
    setIsPostPengeluaranBarangMsg,
    // data table
    setIsdataBarang,
    setIsdataBarangLoading,
} = itemSlice.actions;

export const PostPenerimaanBarang = (input: PenerimaanBarang) => async (dispatch: Dispatch): Promise<void> => {
    try {
        dispatch(setIsPostPenerimaanBarangLoading(true))
        await axios.post<ResponseMessageOnly>(`${BASE_URL}/penerimaan-barang`, input);
        dispatch(setIsPostPenerimaanBarangSuccess(true))
    } catch (error) {
        dispatch(setIsPostPenerimaanBarangSuccess(false))
        dispatch(setIsPostPenerimaanBarangError(true))
        handleError(error, dispatch, setIsPostPenerimaanBarangMsg);
        throw error
    } finally {
        dispatch(setIsPostPenerimaanBarangLoading(false))
    }
};

export const PostPengeluaranBarang = (input: PengeluaranBarang) => async (dispatch: Dispatch): Promise<void> => {
    try {
        dispatch(setIsPostPengeluaranBarangLoading(true))
        await axios.post<ResponseMessageOnly>(`${BASE_URL}/pengeluaran-barang`, input);
        dispatch(setIsPostPengeluaranBarangSuccess(true))
    } catch (error) {
        dispatch(setIsPostPengeluaranBarangSuccess(false))
        dispatch(setIsPostPengeluaranBarangError(true))
        handleError(error, dispatch, setIsPostPengeluaranBarangMsg);
        throw error
    } finally {
        dispatch(setIsPostPengeluaranBarangLoading(false))
    }
};

export const GetReport = () => async (dispatch: Dispatch): Promise<void> => {
    try {
        dispatch(setIsdataBarangLoading(true))
        const response = await axios.get(`${BASE_URL}/report`);
        dispatch(setIsdataBarang(response.data.data))
    } catch (error) {
        throw error
    } finally {
        dispatch(setIsdataBarangLoading(false))
    }
};

export default itemSlice.reducer;
