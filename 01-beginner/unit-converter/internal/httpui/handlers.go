package httpui

import (
	"errors"
	"html/template"
	"math"
	"net/http"
	"strconv"

	"go-backend-labs/01-beginner/unit-converter/internal/convert"
)

type Handler struct {
	tpl *template.Template
}

type UnitOption struct {
	Key   string
	Label string
}

type PageData struct {
	Title  string
	Action string

	Units []UnitOption

	// form state
	Value string
	From  string
	To    string

	// output state
	HasResult bool
	Result    string
	Error     string
	Hint      string
}

func NewHandler() *Handler {
	tpl := template.Must(template.ParseGlob("web/templates/*.html"))
	return &Handler{tpl: tpl}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/length", http.StatusFound)
	})

	mux.HandleFunc("GET /length", h.lengthGet)
	mux.HandleFunc("POST /length", h.lengthPost)

	mux.HandleFunc("GET /weight", h.weightGet)
	mux.HandleFunc("POST /weight", h.weightPost)

	mux.HandleFunc("GET /temperature", h.temperatureGet)
	mux.HandleFunc("POST /temperature", h.temperaturePost)
}

func (h *Handler) render(w http.ResponseWriter, data PageData) {
	if err := h.tpl.ExecuteTemplate(w, "page", data); err != nil {
		http.Error(w, "template render error", http.StatusInternalServerError)
	}
}

func parseFloatField(s string) (float64, error) {
	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return 0, err
	}
	if math.IsNaN(v) || math.IsInf(v, 0) {
		return 0, convert.ErrInvalidValue
	}
	return v, nil
}

func formatResult(v float64) string {
	// Round to 6 decimal places for UI output
	const p = 1e6
	v = math.Round(v*p) / p
	// 'f' + -1 => no trailing zeros
	return strconv.FormatFloat(v, 'f', -1, 64)
}

/* -------------------- Length -------------------- */

var lengthUnits = []UnitOption{
	{"mm", "Millimeter (mm)"},
	{"cm", "Centimeter (cm)"},
	{"m", "Meter (m)"},
	{"km", "Kilometer (km)"},
	{"in", "Inch (in)"},
	{"ft", "Foot (ft)"},
	{"yd", "Yard (yd)"},
	{"mi", "Mile (mi)"},
}

func (h *Handler) lengthGet(w http.ResponseWriter, r *http.Request) {
	h.render(w, PageData{
		Title:  "Length converter",
		Action: "/length",
		Units:  lengthUnits,
		Value:  "1",
		From:   "m",
		To:     "cm",
		Hint:   "Negative values are not allowed for length.",
	})
}

func (h *Handler) lengthPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	data := PageData{
		Title:  "Length converter",
		Action: "/length",
		Units:  lengthUnits,
		Value:  r.FormValue("value"),
		From:   r.FormValue("from"),
		To:     r.FormValue("to"),
		Hint:   "Negative values are not allowed for length.",
	}

	value, err := parseFloatField(data.Value)
	if err != nil {
		data.Error = "Please enter a valid number."
		h.render(w, data)
		return
	}

	out, err := convert.Length(value, data.From, data.To)
	if err != nil {
		if errors.Is(err, convert.ErrNegativeValue) {
			data.Error = "Negative values are not allowed for length."
		} else if errors.Is(err, convert.ErrUnknownUnit) {
			data.Error = "Unknown unit selected."
		} else {
			data.Error = "Conversion error."
		}
		h.render(w, data)
		return
	}

	data.HasResult = true
	data.Result = formatResult(out)
	h.render(w, data)
}

/* -------------------- Weight -------------------- */

var weightUnits = []UnitOption{
	{"mg", "Milligram (mg)"},
	{"g", "Gram (g)"},
	{"kg", "Kilogram (kg)"},
	{"oz", "Ounce (oz)"},
	{"lb", "Pound (lb)"},
}

func (h *Handler) weightGet(w http.ResponseWriter, r *http.Request) {
	h.render(w, PageData{
		Title:  "Weight converter",
		Action: "/weight",
		Units:  weightUnits,
		Value:  "1",
		From:   "kg",
		To:     "g",
		Hint:   "Negative values are not allowed for weight.",
	})
}

func (h *Handler) weightPost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	data := PageData{
		Title:  "Weight converter",
		Action: "/weight",
		Units:  weightUnits,
		Value:  r.FormValue("value"),
		From:   r.FormValue("from"),
		To:     r.FormValue("to"),
		Hint:   "Negative values are not allowed for weight.",
	}

	value, err := parseFloatField(data.Value)
	if err != nil {
		data.Error = "Please enter a valid number."
		h.render(w, data)
		return
	}

	out, err := convert.Weight(value, data.From, data.To)
	if err != nil {
		if errors.Is(err, convert.ErrNegativeValue) {
			data.Error = "Negative values are not allowed for weight."
		} else if errors.Is(err, convert.ErrUnknownUnit) {
			data.Error = "Unknown unit selected."
		} else {
			data.Error = "Conversion error."
		}
		h.render(w, data)
		return
	}

	data.HasResult = true
	data.Result = formatResult(out)
	h.render(w, data)
}

/* -------------------- Temperature -------------------- */

var tempUnits = []UnitOption{
	{"C", "Celsius (C)"},
	{"F", "Fahrenheit (F)"},
	{"K", "Kelvin (K)"},
}

func (h *Handler) temperatureGet(w http.ResponseWriter, r *http.Request) {
	h.render(w, PageData{
		Title:  "Temperature converter",
		Action: "/temperature",
		Units:  tempUnits,
		Value:  "0",
		From:   "C",
		To:     "F",
		Hint:   "Negative values are allowed for temperature.",
	})
}

func (h *Handler) temperaturePost(w http.ResponseWriter, r *http.Request) {
	_ = r.ParseForm()

	data := PageData{
		Title:  "Temperature converter",
		Action: "/temperature",
		Units:  tempUnits,
		Value:  r.FormValue("value"),
		From:   r.FormValue("from"),
		To:     r.FormValue("to"),
		Hint:   "Negative values are allowed for temperature.",
	}

	value, err := parseFloatField(data.Value)
	if err != nil {
		data.Error = "Please enter a valid number."
		h.render(w, data)
		return
	}

	out, err := convert.Temperature(value, data.From, data.To)
	if err != nil {
		if errors.Is(err, convert.ErrUnknownUnit) {
			data.Error = "Unknown unit selected."
		} else {
			data.Error = "Conversion error."
		}
		h.render(w, data)
		return
	}

	data.HasResult = true
	data.Result = formatResult(out)
	h.render(w, data)
}
