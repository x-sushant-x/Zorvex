import '@fortawesome/fontawesome-free/css/all.css';


export function ServiceCard() {
    return (
        <>
            <div className="p-4 rounded-xl w-3/12 border-2 border-slate-300">
                <div className="flex justify-between items-center mb-4">
                    <h1 className="text-xl font-medium">Edit Profile Service</h1>
                    <div className="flex items-center">
                        <div className="bg-red-500 w-8 h-8 rounded-full flex items-center justify-center">
                            <i className="fas fa-times text-white"></i>
                        </div>
                    </div>
                </div>
                <p className="mb-2 text-slate-800">
                    Protocol: https
                </p>
                <p className="text-slate-800">
                    IP Address: 192.124.13.43
                </p>
                <div className="flex items-center justify-between text-slate-800">
                    Port: 8080
                    <button className="bg-slate-950 text-white text-sm py-2 px-4 rounded">
                        Edit
                    </button>
                </div>
            </div>

        </>
    )
}