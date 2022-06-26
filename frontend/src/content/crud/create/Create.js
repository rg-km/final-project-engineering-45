import React, { useState } from "react";
import Navbar from "../../../component/navbar/navbar";
import Footer from "../../../component/footer/footer";
import axios from "axios";
import "./Create.css";

const Create = () => {
    const [prodiname, setProdiname] = useState("");
    const [fakultasname, setFakultasname] = useState("");
    const [deskripsi, setDeskripsi] = useState("");
    const [karakter, setKarakter] = useState("");
    const [matakuliah, setMatakuliah] = useState("");
    const [prospek, setProspek] = useState("");

    const sendData = () => {
        axios
            .post("http://localhost:8080/api/prodi/add", {
                prodi_name: prodiname,
                fakultas_name: fakultasname,
                deskripsi,
                karakter,
                mata_kuliah: matakuliah,
                prospek,
            })
            .then(() => {
                console.log("berhasil");
            });
    };

    return (
        <>
            <Navbar />
            <div className="create-container">
                <div className="header-create">
                    <h1>Form Menambahkan Program Studi</h1>
                    <p>Silahkan tambahkan program studi baru</p>
                </div>
                <div className="form-create">
                    <label htmlFor="prodiname" className="form-label">
                        Prodi Name :{" "}
                    </label>
                    <input
                        type="text"
                        name="prodiname"
                        onChange={(e) => setProdiname(e.target.value)}
                        placeholder="Prodi Name"
                        className="form-control"
                    />
                </div>
                <div className="form-create">
                    <label htmlFor="fakultasname" className="form-label">
                        Fakultas Name :{" "}
                    </label>
                    <input
                        type="text"
                        name="fakultasname"
                        onChange={(e) => setFakultasname(e.target.value)}
                        placeholder="Fakultas Name"
                        className="form-control"
                    />
                </div>
                <div className="form-create">
                    <label htmlFor="deskripsi" className="form-label">
                        Deskripsi :{" "}
                    </label>
                    <input
                        type="text"
                        name="deskripsi"
                        onChange={(e) => setDeskripsi(e.target.value)}
                        placeholder="Deskripsi"
                        className="form-control"
                    />
                </div>
                <div className="form-create">
                    <label htmlFor="karakter" className="form-label">
                        Karakter :{" "}
                    </label>
                    <input
                        type="text"
                        name="karakter"
                        onChange={(e) => setKarakter(e.target.value)}
                        placeholder="Karakter"
                        className="form-control"
                    />
                </div>
                <div className="form-create">
                    <label htmlFor="matakuliah" className="form-label">
                        Mata Kuliah :{" "}
                    </label>
                    <input
                        type="text"
                        name="matakuliah"
                        onChange={(e) => setMatakuliah(e.target.value)}
                        placeholder="Mata Kuliah"
                        className="form-control"
                    />
                </div>
                <div className="form-create">
                    <label htmlFor="prospek" className="form-label">
                        Prospek Kerja:{" "}
                    </label>
                    <input
                        type="text"
                        name="prospek"
                        onChange={(e) => setProspek(e.target.value)}
                        placeholder="Prospek Kerja"
                        className="form-control"
                    />
                </div>
                <button type="submit" onClick={sendData} className="btn-create">
                    Submit
                </button>
            </div>
            <Footer />
        </>
    );
};

export default Create;
