import { AddPenerimaan } from "@/components/ui/add-penerimaan";
import { AddPengeluaran } from "@/components/ui/add-pengeluaran";
import TableBarang from "@/components/ui/TableBarang";

const Home = () => {
  return (
    <>
      <div className="flex flex-col justify-center item-center p-10">
        <div className="flex gap-3">
          <AddPenerimaan />
          <AddPengeluaran />
        </div>
        <TableBarang />
      </div>
    </>
  );
};

export default Home;
