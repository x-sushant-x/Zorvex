function AppBar() {
    return (
        <>
            <header className=" py-4 px-12">
                <div className="mx-auto flex justify-between items-center">
                    <div className="flex items-center">
                        <h1 className="text-xl font-bold">Zorvex</h1>
                    </div>
                    <div className="flex items-center">
                        <nav className="flex space-x-12 pr-12">
                            <a href="#">LinkedIn</a>
                            <a href="#">My Blog</a>
                            <a href="#">GitHub</a>
                        </nav>
                        <input type="text" placeholder="Search Services" className="border border-black rounded-lg py-2 pl-4 pr-4 w-60" />

                    </div>
                </div>
            </header>
            <div className="w-screen h-[1.2px] bg-slate-400"></div>
        </>
    );
}

export default AppBar;
