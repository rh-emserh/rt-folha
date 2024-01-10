// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/unidade": {
            "post": {
                "description": "Returns the unit",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "summary": "Returns the specific unit with information for the pdt",
                "parameters": [
                    {
                        "description": "Unidade",
                        "name": "unidade",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "type": "string"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.PDT"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.PDT": {
            "type": "object",
            "properties": {
                "adicional_noturno_folha": {
                    "type": "number"
                },
                "auxilio_alimentacao": {
                    "type": "number"
                },
                "dsr_folha": {
                    "description": "Adicional_noturnoo int ` + "`" + `json:\"adicional_noturno_folha\"` + "`" + `",
                    "type": "number"
                },
                "encargos_folha": {
                    "type": "number"
                },
                "gratificacao_folha": {
                    "type": "number"
                },
                "hora_extras": {
                    "type": "number"
                },
                "insalubridade_folha": {
                    "type": "number"
                },
                "periculosidade_folha": {
                    "type": "number"
                },
                "rh": {
                    "type": "integer"
                },
                "salario_folha": {
                    "type": "number"
                },
                "setor": {
                    "type": "string"
                },
                "total_anual_folha": {
                    "type": "number"
                },
                "total_mensal_folha": {
                    "description": "vale transporte",
                    "type": "number"
                },
                "total_salario_folha": {
                    "type": "number"
                },
                "unidade": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API FOLHA DE PAGAMENTO",
	Description:      "API for crud operations on users",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
