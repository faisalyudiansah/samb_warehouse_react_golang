import React, { useState, useEffect } from "react";
import { Button } from "@/components/ui/button";
import {
    Dialog,
    DialogContent,
    DialogDescription,
    DialogFooter,
    DialogHeader,
    DialogTitle,
    DialogTrigger,
} from "@/components/ui/dialog";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";
import { useDispatch, useSelector } from "react-redux";
import { AppDispatch, RootState } from "@/stores";
import { PengeluaranBarang } from "@/@types/request";
import Swal from "sweetalert2";
import { GetReport, PostPengeluaranBarang } from "@/stores/slices/itemSlices"; // Adjust API slice

export function AddPengeluaran() {
    const dispatch = useDispatch<AppDispatch>();
    const { isPostPengeluaranBarangSuccess, isPostPengeluaranBarangError, isPostPengeluaranBarangMsg } = useSelector(
        (state: RootState) => state.itemSlice
    );
    const [isDialogOpen, setIsDialogOpen] = useState(false);
    const [input, setInput] = useState<PengeluaranBarang>({
        whs_idf: 0,
        trx_out_date: "",
        trx_out_supp_idf: 0,
        trx_out_notes: "",
        trx_out_dproduct_idf: 0,
        trx_out_dqty_dus: 0,
        trx_out_dqty_pcs: 0,
    });
    const [errorMsg, setErrorMsg] = useState<string>("");

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;
        setInput({
            ...input,
            [name]: name === "whs_idf" || name === "trx_out_supp_idf" || name === "trx_out_dproduct_idf" || name === "trx_out_dqty_dus" || name === "trx_out_dqty_pcs"
                ? Number(value)
                : value,
        });
    };

    const handleSubmit = () => {
        dispatch(PostPengeluaranBarang(input)); 
    };

    useEffect(() => {
        if (isPostPengeluaranBarangSuccess) {
            const fetch = async () => {
                dispatch(GetReport());
            };
            setErrorMsg("");
            fetch();
            setIsDialogOpen(false);
            setInput({
                whs_idf: 0,
                trx_out_date: "",
                trx_out_supp_idf: 0,
                trx_out_notes: "",
                trx_out_dproduct_idf: 0,
                trx_out_dqty_dus: 0,
                trx_out_dqty_pcs: 0,
            });
            Swal.fire({
                title: "Success",
                text: "Pengeluaran Barang berhasil ditambahkan!",
                icon: "success",
                confirmButtonText: "OK",
            });
        } else if (isPostPengeluaranBarangError) {
            setErrorMsg(isPostPengeluaranBarangMsg || "Terjadi kesalahan, silakan coba lagi.");
        }
    }, [isPostPengeluaranBarangSuccess, isPostPengeluaranBarangError, isPostPengeluaranBarangMsg]);

    return (
        <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
            <DialogTrigger asChild>
                <Button variant="outline">Add Pengeluaran</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[725px]">
                <DialogHeader>
                    <DialogTitle>Add Pengeluaran</DialogTitle>
                    <DialogDescription>
                        Tambahkan barang baru pada pengeluaran
                    </DialogDescription>
                    {errorMsg && (
                        <div className="bg-gray-200 p-2 flex items-center justify-center rounded-xl">
                            <p className="text-red-500 text-left">{errorMsg}</p>
                        </div>
                    )}
                </DialogHeader>
                <div className="grid gap-4 py-4">
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="whs_idf" className="text-right">
                            Warehouse ID
                        </Label>
                        <Input
                            id="whs_idf"
                            name="whs_idf"
                            type="number"
                            value={input.whs_idf}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_out_date" className="text-right">
                            Transaction Date
                        </Label>
                        <Input
                            id="trx_out_date"
                            name="trx_out_date"
                            type="date"
                            value={input.trx_out_date}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_out_supp_idf" className="text-right">
                            Supplier ID
                        </Label>
                        <Input
                            id="trx_out_supp_idf"
                            name="trx_out_supp_idf"
                            type="number"
                            value={input.trx_out_supp_idf}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_out_dproduct_idf" className="text-right">
                            Product ID
                        </Label>
                        <Input
                            id="trx_out_dproduct_idf"
                            name="trx_out_dproduct_idf"
                            type="number"
                            value={input.trx_out_dproduct_idf}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_out_dqty_dus" className="text-right">
                            Quantity Dus
                        </Label>
                        <Input
                            id="trx_out_dqty_dus"
                            name="trx_out_dqty_dus"
                            type="number"
                            value={input.trx_out_dqty_dus}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_out_dqty_pcs" className="text-right">
                            Quantity Pcs
                        </Label>
                        <Input
                            id="trx_out_dqty_pcs"
                            name="trx_out_dqty_pcs"
                            type="number"
                            value={input.trx_out_dqty_pcs}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_out_notes" className="text-right">
                            Notes
                        </Label>
                        <textarea
                            id="trx_out_notes"
                            name="trx_out_notes"
                            value={input.trx_out_notes}
                            onChange={handleChange}
                            className="col-span-3 border rounded dark:text-black px-3 py-2"
                            rows={4}
                        />
                    </div>
                </div>
                <DialogFooter>
                    <Button
                        type="submit"
                        onClick={handleSubmit}
                        disabled={!input.whs_idf || !input.trx_out_date || !input.trx_out_supp_idf || !input.trx_out_dproduct_idf || !input.trx_out_dqty_dus || !input.trx_out_dqty_pcs}
                    >
                        Save changes
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    );
}
