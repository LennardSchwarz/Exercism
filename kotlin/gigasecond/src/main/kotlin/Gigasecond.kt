import java.time.LocalDate
import java.time.LocalDateTime

class Gigasecond(baseDatetime: LocalDateTime) {
    private val gigaseconds: Long = 1_000_000_000
    val date: LocalDateTime = baseDatetime.plusSeconds(gigaseconds)

    constructor(baseDate: LocalDate) : this(baseDate.atStartOfDay())
}