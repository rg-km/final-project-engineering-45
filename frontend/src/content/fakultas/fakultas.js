import React from 'react'
import {Link} from 'react-router-dom';
import Navbar from '../../component/navbar/navbar';
import Footer from '../../component/footer/footer';
import Button from '../../component/button/button';
import {MenuFakultas} from './menuFakultas';
import './fakultas.css';

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
                                    <Link to={item.url} className="more-content">
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
        </div>
    );
};

export default Fakultas;
