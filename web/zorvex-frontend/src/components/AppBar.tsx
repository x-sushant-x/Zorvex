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
                    </div>
                </div>
            </header>
            <div className="w-screen h-[1.2px] bg-slate-400"></div>
        </>
    );
}

export default AppBar;
