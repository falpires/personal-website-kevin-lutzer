import { HttpErrorResponse } from "@angular/common/http";
import { Observable, throwError } from "rxjs";

export interface ApiResponse<T>{
  success: string;
  result?: T;
}


export class ApiError {
  constructor(public error: string, public status?: number) {}
}

export function defaultErrorHandler(error: HttpErrorResponse | string | ApiError): Observable<any> {
  console.log(error);

  if (typeof (<any> error)?.error == 'string')
    return throwError((<any>error)?.error);
  else if(error instanceof ApiError || typeof error == 'string')
    return throwError(error);
  else if (error.error instanceof ErrorEvent) {
    return throwError(error.error.message);
  }

  console.error(`Backend returned code ${error.status}, body was: ${error.error}`);
  const resp = new ApiError(error?.error?.error, error?.status);
  return throwError(resp);
}
