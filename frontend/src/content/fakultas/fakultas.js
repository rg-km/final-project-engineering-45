import React, {useState, useEffect} from "react";
import { Outlet, Link } from "react-router-dom";
import Navbar from "../../component/navbar/navbar";
import Footer from "../../component/footer/footer";
import Button from "../../component/button/button";
import { MenuFakultas } from "./menuFakultas";
// import axios from 'axios';
import "./fakultas.css";

const Fakultas = () => {
    return (
        <div>
            <Navbar />
            <div class="FakultasSection">
                <div className="heading">
                    <h1>List Fakultas</h1>
                    <p>Simak fakultas yang sesuai dengan passion kalian</p>
                </div>
                <div className="MenuFakultas">
                    <ul className="MenuFakultas-list">
                        {MenuFakultas.map((item, index) => (
                            <li key={index} className="fakultas-list">
                                <h1 className="header">RightWay</h1>
                                <div className="link-text">
                                    <Link className={item.cName} to={item.url}>
                                        {item.title}
                                    </Link>
                                    <Link
                                        to={item.url}
                                        className="more-content"
                                    >
                                        Lihat selengkapnya
                                    </Link>
                                </div>
                            </li>
                        ))}
                    </ul>
                </div>
                <div className="btn-see-more">
                    <Link to="/sign-up">
                        <Button
                            buttonStyle="btn--primary"
                            children="Lihat fakultas lainnya"
                        />
                    </Link>
                </div>
            </div>
            <Footer />
            <Outlet />
        </div>
    );
};

// const Fakultas = () => {
//     const [data, setData] = useState([]);
//     const [loading, setLoading] = useState(false);
//     const [error, setError] = useState(false);

//     useEffect(() => {
//         async function getData() {
//             setLoading(true);
//             setError(false);
//             try {
//                 const result = await axios.get('http://localhost:8080/api/fakultas/list');
//                 setData(result.data);
//             }
//             catch (err) {
//                 setError(true);
//             }
//             setLoading(false);
//         }
//         getData();
//     }, [])

//     return (
//         <div>
//             <Navbar />
//             <div class="FakultasSection">
//                 <div className="heading">
//                     <h1>List Fakultas</h1>
//                     <p>Simak fakultas yang sesuai dengan passion kalian</p>
//                 </div>
//                 <div>
//                     {loading && <p>Loading...</p>}
//                     {error && <p>Error</p>}
//                     {data.map((item, index) => (
//                         <ul key={index}>
//                             <li>{item.fakultas_name}</li>
//                         </ul>
//                     ))}
//                 </div>
//                 <div className="btn-see-more">
//                     <Link to="/sign-up">
//                         <Button
//                             buttonStyle="btn--primary"
//                             children="Lihat fakultas lainnya"
//                         />
//                     </Link>
//                 </div>
//             </div>
//             <Footer />
//             <Outlet />
//         </div>
//     );
// };

export default Fakultas;
