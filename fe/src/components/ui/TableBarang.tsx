import { useDispatch, useSelector } from "react-redux"
import { AppDispatch, RootState } from "@/stores"
import { useEffect } from "react"
import { GetReport } from "@/stores/slices/itemSlices"

const TableBarang = () => {
    const dispatch = useDispatch<AppDispatch>()
    const { dataBarang, dataBarangLoading } = useSelector(
        (state: RootState) => state.itemSlice
    )

    useEffect(() => {
        const fetch = async () => {
            dispatch(GetReport())
        }
        fetch()
    }, [dispatch])

    return (
        <div>
            <h2 className="text-lg font-bold mb-4">Laporan Barang</h2>
            {dataBarangLoading ? (
                <p>Loading...</p>
            ) : (
                <table className="min-w-full bg-white border">
                    <thead>
                        <tr>
                            <th className="px-4 py-2 border dark:text-black">Warehouse</th>
                            <th className="px-4 py-2 border dark:text-black">Product</th>
                            <th className="px-4 py-2 border dark:text-black">Qty Dus</th>
                            <th className="px-4 py-2 border dark:text-black">Qty Pcs</th>
                            <th className="px-4 py-2 border dark:text-black">Type</th>
                        </tr>
                    </thead>
                    <tbody>
                        {dataBarang.map((item, index) => (
                            <tr key={index} className={`${item.type === "Penerimaan" ? "bg-gray-200" : ""}`}>
                                <td className="px-4 py-2 border dark:text-black">{item.warehouse}</td>
                                <td className="px-4 py-2 border dark:text-black">{item.product}</td>
                                <td className="px-4 py-2 border dark:text-black">{item.qty_dus}</td>
                                <td className="px-4 py-2 border dark:text-black">{item.qty_pcs}</td>
                                <td className="px-4 py-2 border dark:text-black">{item.type}</td>
                            </tr>
                        ))}
                    </tbody>
                </table>
            )}
        </div>
    )
}

export default TableBarang
