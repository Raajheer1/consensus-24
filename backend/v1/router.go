package v1

import (
	"github.com/go-chi/chi/v5"
	"steller-api/v1/account"
	"steller-api/v1/feed"
	"steller-api/v1/gov"
	"steller-api/v1/gov/votes"
	"steller-api/v1/loan"
	loan_lender "steller-api/v1/loan-lender"
)

func Router(r chi.Router) {
	r.Route("/v1", func(r chi.Router) {
		r.Route("/account", func(r chi.Router) {
			r.Get("/{id}", account.GetAccount)
			r.Get("/{id}/loan_lendings", loan_lender.GetByLenderID)
		})

		r.Route("/loan", func(r chi.Router) {
			r.Post("/", loan.CreateLoan)
			r.Get("/", loan.GetLoans)
			r.Get("/funding", loan.GetInactiveLoans)
			r.Route("/{loanId}", func(r chi.Router) {
				r.Route("/gv", func(r chi.Router) {
					r.Get("/", gov.GetGovVotes)
					r.Post("/", gov.CreateGovVote)
					r.Route("/{govId}", func(r chi.Router) {
						r.Get("/votes", votes.GetVotes)
						r.Post("/vote", votes.Vote)
					})
				})
				r.Route("/ll", func(r chi.Router) {
					r.Get("/", loan_lender.GetLoanLenders)
					r.Post("/", loan_lender.CreateLoanLender)
				})
			})
		})

		r.Route("/feed/{feedType}/{id}", func(r chi.Router) {
			r.Get("/", feed.GetFeed)
			r.Post("/", feed.CreateFeed)
		})
	})
}
