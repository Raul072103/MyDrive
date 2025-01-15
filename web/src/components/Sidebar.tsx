import "@styles/global.css";
import {Link} from "react-router-dom";

const Sidebar = () => {
    const menuItems = [
        { name: "Home", link: "/home" },
        { name: "My Files", link: "/home/myfiles" },
        { name: "Starred", link: "/home/starred" },
        { name: "Trash", link: "/home/trash" },
    ];

    return (
        <div className="sidebar">
            <ul>
                {menuItems.map((item, index) => (
                    <li key={index}>
                        <Link to={item.link}>
                            <i className=""></i> {item.name}
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Sidebar;
