/// <reference path="./global.d.ts" />
// @ts-check

/**
 * Creates a new visitor.
 *
 * @param {string} name
 * @param {number} age
 * @param {string} ticketId
 * @returns {Visitor} the visitor that was created
 */
export function createVisitor(name, age, ticketId) {
 return {name, age, ticketId} 
}

/**
 * Revokes a ticket for a visitor.
 *
 * @param {Visitor} visitor the visitor with an active ticket
 * @returns {Visitor} the visitor without a ticket
 */
export function revokeTicket(visitor) {
  visitor.ticketId = null
  return visitor
}

/**
 * Determines the status a ticket has in the ticket tracking object.
 *
 * @param {Record<string, string|null>} tickets
 * @param {string} ticketId
 * @returns {string} ticket status
 */
export function ticketStatus(tickets, ticketId) {
  const visitor = tickets[ticketId]
  if (typeof visitor === 'string') {
    return 'sold to ' + visitor;
  }
  else if (visitor === null) {
    return 'not sold'
  } else {
    return 'unknown ticket id'
  }
}

/**
 * Determines the status a ticket has in the ticket tracking object
 * and returns a simplified status message.
 *
 * @param {Record<string, string|null>} tickets
 * @param {string} ticketId
 * @returns {string} ticket status
 */
export function simpleTicketStatus(tickets, ticketId) {
// Check if the ticketId exists in the tickets object and is not null
if (ticketId in tickets && tickets[ticketId] !== null) {
  // Return the visitor name or empty string (for "strange" name values)
  return tickets[ticketId] || "Ticket is valid but no name provided";
} else {
  // Return invalid ticket message for tickets not found or with null values
  return 'invalid ticket !!!';
}
}

/**
 * Determines the version of the GTC that was signed by the visitor.
 *
 * @param {VisitorWithGtc} visitor
 * @returns {string | undefined} version
 */
export function gtcVersion(visitor) {
  return visitor.gtc?.version
}
