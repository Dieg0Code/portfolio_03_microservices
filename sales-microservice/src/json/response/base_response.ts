interface BaseResponse<T> {
    code: number;
    status: string;
    msg: string;
    data: T;
}