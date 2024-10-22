import { Outlet } from 'react-router-dom'
import { Navbar } from '../ui/navbar'

const MainLayout = () => {
  return (
    <>
      <Navbar />
      <Outlet />
    </>
  )
}

export default MainLayout