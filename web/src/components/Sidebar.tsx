import "./Sidebar.css";
import {Link} from "react-router-dom";

const Sidebar = () => {
    const menuItems = [
        { name: "Home", link: "/home", icon: "/assets/icons/home.svg" },
        { name: "My Files", link: "/home/myfiles", icon: "/assets/icons/folder.svg" },
        { name: "Starred", link: "/home/starred", icon: "/assets/icons/star.svg" },
        { name: "Trash", link: "/home/trash", icon: "/assets/icons/trash.svg" },
    ];

    return (
        <div className="sidebar">
            <ul>
                {menuItems.map((item, index) => (
                    <li key={index}>
                        <Link to={item.link}>
                            <img src={item.icon} alt={`${item.name} icon`} className="icon"/>
                            {item.name}
                        </Link>
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Sidebar;
