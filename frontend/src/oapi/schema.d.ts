/**
 * This file was auto-generated by openapi-typescript.
 * Do not make direct changes to the file.
 */


export interface paths {
  "/hello/{name}": {
    /**
     * Greeter
     * @description Greeter greets you.
     */
    get: operations["main"];
  };
}

export type webhooks = Record<string, never>;

export interface components {
  schemas: {
    HelloOutput: {
      message?: string;
    };
    RestErrResponse: {
      /** @description Application-specific error code. */
      code?: number;
      /** @description Application context. */
      context?: {
        [key: string]: unknown;
      };
      /** @description Error message. */
      error?: string;
      /** @description Status text. */
      status?: string;
    };
  };
  responses: never;
  parameters: never;
  requestBodies: never;
  headers: never;
  pathItems: never;
}

export type $defs = Record<string, never>;

export type external = Record<string, never>;

export interface operations {

  /**
   * Greeter
   * @description Greeter greets you.
   */
  main: {
    parameters: {
      query?: {
        locale?: "en-US";
      };
      path: {
        name: string;
      };
    };
    responses: {
      /** @description OK */
      200: {
        headers: {
          "X-Now"?: string;
        };
        content: {
          "application/json": components["schemas"]["HelloOutput"];
        };
      };
      /** @description Bad Request */
      400: {
        content: {
          "application/json": components["schemas"]["RestErrResponse"];
        };
      };
    };
  };
}
