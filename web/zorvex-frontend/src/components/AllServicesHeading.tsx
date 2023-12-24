function AllServicesHeading() {
    return (
        <>
            <div className="flex justify-center h-screen mt-10">
                <div className="container w-5/6 h-20 flex justify-between">

                    <div>
                        <h1 className="text-3xl font-medium mb-2">All Services</h1>
                        <p className="text-slate-700">Services that are currently registered</p>
                    </div>

                    <div className="flex justify-between items-center px-4 ">
                        <button className="bg-slate-950 text-white text-sm py-2 px-4 rounded">
                            Add New
                        </button>
                    </div>

                </div>
            </div>

        </>
    )
}

export default AllServicesHeading