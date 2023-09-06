package domain

// Определение структуры Event для представления событий
type Event struct {
    ID      int    `json:"id"`
    Title   string `json:"title"`
    Date    string `json:"date"`
    Details string `json:"details"`
}
