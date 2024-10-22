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
import { PenerimaanBarang } from "@/@types/request";
import Swal from "sweetalert2";
import { GetReport, PostPenerimaanBarang } from "@/stores/slices/itemSlices";

export function AddPenerimaan() {
    const dispatch = useDispatch<AppDispatch>();
    const { isPostPenerimaanBarangSuccess, isPostPenerimaanBarangError, isPostPenerimaanBarangMsg } = useSelector(
        (state: RootState) => state.itemSlice
    );
    const [isDialogOpen, setIsDialogOpen] = useState(false);
    const [input, setInput] = useState<PenerimaanBarang>({
        whs_idf: 0,
        trx_in_date: "",
        trx_in_supp_idf: 0,
        trx_in_notes: "",
        trx_in_dproduct_idf: 0,
        trx_in_dqty_dus: 0,
        trx_in_dqty_pcs: 0,
    });
    const [errorMsg, setErrorMsg] = useState<string>("");

    const handleChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        const { name, value } = e.target;

        setInput({
            ...input,
            [name]: name === "whs_idf" || name === "trx_in_supp_idf" || name === "trx_in_dproduct_idf" || name === "trx_in_dqty_dus" || name === "trx_in_dqty_pcs"
                ? Number(value)
                : value,
        });
    };

    const handleSubmit = () => {
        dispatch(PostPenerimaanBarang(input));
    };

    useEffect(() => {
        if (isPostPenerimaanBarangSuccess) {
            const fetch = async () => {
                dispatch(GetReport())
            }
            setErrorMsg("");
            setIsDialogOpen(false);
            fetch()
            setInput({
                whs_idf: 0,
                trx_in_date: "",
                trx_in_supp_idf: 0,
                trx_in_notes: "",
                trx_in_dproduct_idf: 0,
                trx_in_dqty_dus: 0,
                trx_in_dqty_pcs: 0,
            });
            Swal.fire({
                title: "Success",
                text: "Penerimaan Barang berhasil ditambahkan!",
                icon: "success",
            });
        } else if (isPostPenerimaanBarangError) {
            setErrorMsg(isPostPenerimaanBarangMsg || "Terjadi kesalahan, silakan coba lagi.");
        }
    }, [isPostPenerimaanBarangSuccess, isPostPenerimaanBarangError, isPostPenerimaanBarangMsg]);

    return (
        <Dialog open={isDialogOpen} onOpenChange={setIsDialogOpen}>
            <DialogTrigger asChild>
                <Button variant="outline">Add Penerimaan</Button>
            </DialogTrigger>
            <DialogContent className="sm:max-w-[725px]">
                <DialogHeader>
                    <DialogTitle>Add Penerimaan</DialogTitle>
                    <DialogDescription>
                        Tambahkan barang baru pada penerimaan
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
                        <Label htmlFor="trx_in_date" className="text-right">
                            Transaction Date
                        </Label>
                        <Input
                            id="trx_in_date"
                            name="trx_in_date"
                            type="date"
                            value={input.trx_in_date}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_in_supp_idf" className="text-right">
                            Supplier ID
                        </Label>
                        <Input
                            id="trx_in_supp_idf"
                            name="trx_in_supp_idf"
                            type="number"
                            value={input.trx_in_supp_idf}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_in_dproduct_idf" className="text-right">
                            Product ID
                        </Label>
                        <Input
                            id="trx_in_dproduct_idf"
                            name="trx_in_dproduct_idf"
                            type="number"
                            value={input.trx_in_dproduct_idf}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_in_dqty_dus" className="text-right">
                            Quantity Dus
                        </Label>
                        <Input
                            id="trx_in_dqty_dus"
                            name="trx_in_dqty_dus"
                            type="number"
                            value={input.trx_in_dqty_dus}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_in_dqty_pcs" className="text-right">
                            Quantity Pcs
                        </Label>
                        <Input
                            id="trx_in_dqty_pcs"
                            name="trx_in_dqty_pcs"
                            type="number"
                            value={input.trx_in_dqty_pcs}
                            onChange={handleChange}
                            className="col-span-3"
                        />
                    </div>
                    <div className="grid grid-cols-4 items-center gap-4">
                        <Label htmlFor="trx_in_notes" className="text-right">
                            Notes
                        </Label>
                        <textarea
                            id="trx_in_notes"
                            name="trx_in_notes"
                            value={input.trx_in_notes}
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
                        disabled={
                            !input.whs_idf ||
                            !input.trx_in_date ||
                            !input.trx_in_supp_idf ||
                            !input.trx_in_dproduct_idf ||
                            !input.trx_in_dqty_dus ||
                            !input.trx_in_dqty_pcs
                        }
                    >
                        Save changes
                    </Button>
                </DialogFooter>
            </DialogContent>
        </Dialog>
    );
}
