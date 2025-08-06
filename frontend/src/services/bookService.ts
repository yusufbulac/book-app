import axios from "axios";
import {BookResponse} from "@/types/book";

export const fetchBooks = async (): Promise<BookResponse[]> => {
    const response = await axios.get("http://localhost:8080/api/v1/books");
    return response.data.data;
};
