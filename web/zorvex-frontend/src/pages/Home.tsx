import AllServicesHeading from '../components/AllServicesHeading'
import { ServiceCard } from '../components/ServiceCard'
import AppBar from './../components/AppBar'

function Home() {
    return (
        <>
            <AppBar />
            <AllServicesHeading />

            <div className='flex space-x-4 justify-center mt-12'>
                <ServiceCard />
                <ServiceCard />
                <ServiceCard />

            </div>
        </>
    )
}

export default Home